# Ad Hoc
Ad hoc is a indexation-based notes system, based in what wikipedia does.

Every note is related to some other one by indexes, for example:


Destine:
```markdown
# Machine learning uses
[Machine learning] is wide used is health, astronomy, etc.
```

Origin:
```markdown
# Machine learning concept
[Machine learning] comes from Data science, Maths and Statistics fields. Is defined as...
```

## Basic operation

There are three main tables in the database:

### Ideas
The base of the note content. Here, the user enter the title and the content of the note, but, during the user input,
keywords are also requested. These keywords are saved separately in the `Kewords` table.

| id (BIGINT) | title VARCHAR(255)          | content VARCHAR(255)                |
|-------------|-----------------------------|-------------------------------------|
| 1           | my interests                | i like machine learning for biology |
| 2           | available careers           | studies in biology, statistics      |
| 3           | current trend topics fields | machine learning, biology |


### Keywords
After a brief validation and cleaning, unique terms are stored right here.

| id BIGINT | keyword VARCHAR(255) |
|-----------|----------------------|
| 1         | machine learning     |
| 2         | biology              |
| 3         | statistics           |

### Idea_Keywords (many-to-many)

This junction table maps each idea to its associated keywords.
It connects notes to their conceptual terms, enabling graph traversals and weight calculations based on shared terms.

| id BIGINT | idea_id BIGINT | keyword_id BIGINT |
|-----------|----------------|-------------------|
| 1         | 1              | 1                 |
| 2         | 1              | 2                 |
| 3         | 2              | 2                 |
| 4         | 2              | 3                 |
| 5         | 3              | 1                 |
| 6         | 3              | 2                 |

### Edges
The edges of the graph are derived directly from the intersection of keywords between ideas.
Weights are calculated as the total unique keywords count between ideas.
For this example, the idea 1 shares only one keyword with idea 2, but shares 2 keywords with the idea 3, then
the weights for each one will be 1 and 2 respectively.

| id BIGINT | origin_idea_id BIGINT | destine_idea_id BIGINT | weight |
|-----------|-----------------------|------------------------|--------|
| 1         | 1                     | 2                      | 1      |
| 2         | 1                     | 3                      | 2      |
| 3         | 2                     | 1                      | 1      |
| 4         | 2                     | 3                      | 1      |
| 5         | 3                     | 1                      | 2      |
| 6         | 3                     | 2                      | 1      |