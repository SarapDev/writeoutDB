# Write Our DB

Simple project to get understand how DBMS work.

## Roadmap:
-  Start with a simple in-memory key-value store. []
    - can create, delete, modify records []
    - have HTTP access. []
    - store everything in mape []
    - don't let you write too much in parallel (mutex, lock) []
- Extend it to a file-backed database. []
    - File as a backup []
    - File as primary storage []
- Add query parsing and execution. []
    - There is a console access REPL of its own syntax []
- Introduce indexing and concurrency. []
    - Data in files should be structured as a B tree or similar data type. []
- Build a simple replication system. []
