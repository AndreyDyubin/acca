--
-- multi-accounts by user
-- if the user has bonus, credit and other accounts
-- required special format accounts.key for multi-accounts by user

-- required key format (or similar)
-- 'v1.u2.p3.t4.some.fields'
-- v1 - version
-- u2 - userID
-- p3 - purpose account - own funds, etc (for internal)
-- t4 - type account (main, credit, bonus, etc)
-- some.fields - some values

CREATE OR REPLACE FUNCTION acca.ma_get_user_id(key ltree, OUT v varchar)
RETURNS varchar AS $$
    begin
        IF NOT 'ma' @> key THEN
            RAISE EXCEPTION 'Invalid key of account - want prefix "ma", got %', key::text;
        END IF;

        v := cast(subltree(key,1, 2) as varchar);
    end;
$$ language plpgsql;

-- return account type
-- if not exists
-- helper function
CREATE OR REPLACE FUNCTION acca.ma_get_type(key ltree, OUT v varchar)
RETURNS varchar AS $$
    begin
        IF NOT 'ma' @> key THEN
            RAISE EXCEPTION 'Invalid key of account - want prefix "ma", got %', key::text;
        END IF;

        v := cast(subltree(key,3, 4) as varchar);
    end;
$$ language plpgsql;

-- return negative amount if type is credit
-- in other cases return positive amount
--
-- helper function
CREATE OR REPLACE FUNCTION acca.ma_credit_or_debit_by_type(
    _type varchar,
    _balance numeric,
    OUT result numeric
) RETURNS numeric AS $$
    BEGIN
        CASE _type
            WHEN 'credit' THEN
                result = _balance * -1;
            ELSE
                result = _balance;
        END CASE;
    END;
$$ language plpgsql;


CREATE OR REPLACE VIEW acca.ma_accounts AS
    SELECT
        acca.ma_get_user_id(key) AS user_id,
        array_agg(acc_id) AS acc_ids,
        sum(acca.ma_credit_or_debit_by_type(acca.ma_get_type(key), balance)) AS balance_total
    FROM acca.accounts
    WHERE 'ma' @> key
    GROUP BY user_id;

ALTER TABLE acca.balance_changes ADD COLUMN ma_balance jsonb;

CREATE FUNCTION ma_update_balance() RETURNS trigger AS $$
    DECLARE
        _ma_balances jsonb;
        _key ltree;
    BEGIN
        SELECT key INTO _key FROM acca.accounts WHERE acc_id = NEW.acc_id;

        IF NOT 'ma' @> _key THEN
            RETURN NEW;
        END IF;

        select array_to_json(array_agg(row_to_json(tt)))::jsonb INTO _ma_balances
            from (
                SELECT acca.ma_get_type(key) as t, balance as b FROM acca.accounts WHERE subltree(_key, 0, 2) @> key
            ) tt;

        UPDATE acca.balance_changes SET ma_balance = _ma_balances WHERE ch_id = NEW.ch_id;

        RETURN NEW;
    END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_ma_balance_by_bc_trigger AFTER INSERT ON acca.balance_changes
    FOR EACH ROW EXECUTE PROCEDURE ma_update_balance();

-- REPLACE exists view

DROP VIEW acca.recent_activity;
CREATE OR REPLACE VIEW acca.recent_activity AS
    SELECT
        bc.ch_id as id,
        bc.oper_id as oper_id,
        bc.acc_id as acc_id,
        bc.amount as amount,
        bc.balance as balance,
        bc.ma_balance as ma_balance,

        -- operations
        o.tx_id as tx_id,
        o.src_acc_id as src_acc_id,
        o.dst_acc_id as dst_acc_id,
        -- o.type as type,
        o.reason as reason,
        -- o.meta as meta,
        -- o.hold as hold,
        -- o.hold_acc_id as hold_acc_id,
        -- o.status as status,

        -- transactions
        t.reason AS tx_reason,

        -- accounts
        a.key AS acc_key,
         -- currency
        c.curr_id AS acc_curr_id,
        c.key AS acc_curr_key

    FROM acca.balance_changes bc
    LEFT JOIN acca.operations o USING(oper_id)
    LEFT JOIN acca.transactions t USING(tx_id)
    LEFT JOIN acca.accounts a USING(acc_id)
    LEFT JOIN acca.currencies c USING(curr_id)
    ORDER BY id DESC;
