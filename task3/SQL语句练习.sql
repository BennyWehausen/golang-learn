/*假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：*/
//编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
INSERT INTO students (name, age, grade) VALUES ('张三', 20, '三年级');
//编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
SELECT * FROM students WHERE age > 18;
//编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
UPDATE students SET grade = '四年级' WHERE name = '张三';
//编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
DELETE FROM students WHERE age < 15;

/*
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
*/
DELIMITER $$
-- 创建存储过程
CREATE PROCEDURE transfer_funds(
    IN p_from_account_id INT,
    IN p_to_account_id INT
)
BEGIN
    DECLARE from_balance DECIMAL(10,2);
    DECLARE to_balance DECIMAL(10,2);
    DECLARE EXIT HANDLER FOR SQLEXCEPTION ROLLBACK;

START TRANSACTION;

-- 锁定账户 A 和 B 的行（防止并发修改）
SELECT balance INTO from_balance FROM accounts WHERE id = p_from_account_id FOR UPDATE;
SELECT balance INTO to_balance FROM accounts WHERE id = p_to_account_id FOR UPDATE;

-- 检查账户 A 余额是否足够
IF from_balance >= 100 THEN
        -- 扣除账户 A 余额
UPDATE accounts SET balance = balance - 100 WHERE id = p_from_account_id;
-- 增加账户 B 余额
UPDATE accounts SET balance = balance + 100 WHERE id = p_to_account_id;
-- 插入交易记录
INSERT INTO transactions (from_account_id, to_account_id, amount)
VALUES (p_from_account_id, p_to_account_id, 100);
COMMIT;
ELSE
        -- 余额不足，回滚事务
        ROLLBACK;
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = '余额不足，转账失败';
END IF;
END $$

DELIMITER ;