STRUCTURE
$ migrate -source file:///home/henry/work/src/recordkeeping/app/migrations -database postgres://sysdba:orbdba@localhost:5432/recordkeep down

$ migrate -source file:///home/henry/work/src/recordkeeping/app/migrations -database postgres://sysdba:orbdba@localhost:5432/recordkeep up

(Force)
$ migrate -source file:///home/henry/work/src/recordkeeping/app/migrations -database postgres://sysdba:orbdba@localhost:5432/recordkeep force 1519742319