CREATE TABLE posts (
                       PostID INT AUTO_INCREMENT PRIMARY KEY,                -- 帖子ID，自增主键
                       UserID INT NOT NULL,                                  -- 发帖用户ID
                       CommunityID INT NOT NULL,                             -- 所属社区ID（逻辑外键，不定义约束）
                       Title VARCHAR(255) NOT NULL,                          -- 帖子标题
                       Summary VARCHAR(500) NULL,                            -- 帖子摘要（可为空）
                       Content TEXT NOT NULL,                                -- 帖子内容
                       VoteCount INT DEFAULT 0 NOT NULL,                     -- 投票数，默认值为0
                       CreateTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP,       -- 创建时间，默认为当前时间
                       UpdatedTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP
                           ON UPDATE CURRENT_TIMESTAMP                       -- 更新时间，更新时自动修改
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci; -- 使用UTF-8编码，支持中文

