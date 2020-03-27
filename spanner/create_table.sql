CREATE TABLE Books (
  BookId  STRING(36) NOT NULL,
  Name  STRING(1024) NOT NULL,
  Author STRING(255) NOT NULL,
  CreatedAt TIMESTAMP NOT NULL OPTIONS (allow_commit_timestamp=true),
  UpdatedAt TIMESTAMP NOT NULL OPTIONS (allow_commit_timestamp=true),
) PRIMARY KEY (BookId);
