CREATE TABLE community (
                           communityId INT AUTO_INCREMENT PRIMARY KEY,             -- 社区ID，自增主键
                           communityName VARCHAR(128) NOT NULL,                    -- 社区名称，唯一
                           introduction VARCHAR(256) NOT NULL,                     -- 社区简介
                           createTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL, -- 创建时间，默认当前时间
                           updateTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP, -- 更新时间，更新时自动修改
                           CONSTRAINT idx_communityName UNIQUE (communityName)      -- 社区名称唯一约束
)
    COLLATE = utf8mb4_general_ci; -- 使用UTF-8编码，支持中文