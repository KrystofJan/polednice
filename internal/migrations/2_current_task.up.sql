CREATE TABLE current_task (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    current_entry_id INTEGER,
    FOREIGN KEY(current_entry_id) REFERENCES entry(id)
);

CREATE TRIGGER delete_current_task_on_new_insert
BEFORE INSERT ON current_task
FOR EACH ROW
BEGIN
    DELETE FROM current_task;
END;
