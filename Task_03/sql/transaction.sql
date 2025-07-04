BEGIN;

-- 检查账户A的余额是否足够
DECLARE @account_a_balance DECIMAL(18, 2);
DECLARE @account_a_id INT = [账户A的ID]; -- 替换为实际的账户A ID
DECLARE @account_b_id INT = [账户B的ID]; -- 替换为实际的账户B ID
DECLARE @transfer_amount DECIMAL(18, 2) = 100.00;

SELECT @account_a_balance = balance FROM accounts WHERE id = @account_a_id FOR UPDATE;

-- 如果余额不足，回滚事务
IF @account_a_balance < @transfer_amount OR @account_a_balance IS NULL
BEGIN
    ROLLBACK;
    -- 可以在这里添加错误处理逻辑或返回错误信息
    SELECT '转账失败: 余额不足或账户不存在' AS result;
    RETURN;
END

-- 扣除账户A的余额
UPDATE accounts 
SET balance = balance - @transfer_amount 
WHERE id = @account_a_id;

-- 增加账户B的余额
UPDATE accounts 
SET balance = balance + @transfer_amount 
WHERE id = @account_b_id;

-- 记录交易信息到transactions表
INSERT INTO transactions (from_account_id, to_account_id, amount, created_at)
VALUES (@account_a_id, @account_b_id, @transfer_amount, CURRENT_TIMESTAMP);

COMMIT;

SELECT '转账成功' AS result;
