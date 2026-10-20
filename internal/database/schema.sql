-- internal/database/schema.sql

CREATE TABLE ideas (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    content TEXT NOT NULL
);

CREATE TABLE keywords (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    keyword TEXT NOT NULL UNIQUE
);

CREATE TABLE idea_keywords (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    idea_id INTEGER NOT NULL,
    keyword_id INTEGER NOT NULL,
    UNIQUE(idea_id, keyword_id),
    FOREIGN KEY (idea_id) REFERENCES ideas(id) ON DELETE CASCADE,
    FOREIGN KEY (keyword_id) REFERENCES keywords(id) ON DELETE CASCADE
);

CREATE TABLE edges (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    source_idea INTEGER NOT NULL,
    target_idea INTEGER NOT NULL,
    weight INTEGER NOT NULL DEFAULT 1,
    UNIQUE(source_idea, target_idea),
    FOREIGN KEY (source_idea) REFERENCES ideas(id) ON DELETE CASCADE,
    FOREIGN KEY (target_idea) REFERENCES ideas(id) ON DELETE CASCADE
);