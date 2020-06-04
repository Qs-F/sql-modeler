# cmd `sql-modeler`

This cmd simplifies `SQL -> SQLite3 -> sqlboiler -> models` flow.

## Usage

`$ sql-modeler -sql testdata/schema.sql`

If you want to specify the output directly (default is `models` in current directly),

`$ sql-modeler -o /path/to/hogehoge -pkg hogehoge -sql testdata/schema.sql`

You can also define MySQL schema by using https://github.com/dumblob/mysql2sqlite

## Acknowledgements

- `testdata/schema.sql` is retrieved from https://github.com/mattn/sqlboiler-example
- `_example/schema.sql` and `_example/main.go` are retrieved from https://github.com/mattn/sqlboiler-example

## License

MIT License
