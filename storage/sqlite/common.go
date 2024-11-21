package sqlite

const Path_to_file = "output/db-orders.sqlite"

const SchemaSQL = `
CREATE TABLE IF NOT EXISTS last_date (
   id VARCHAR(36),
   table_name VARCHAR(10),
   last_date BITINT(8)
);

CREATE TABLE IF NOT EXISTS exists_pairs (
   id VARCHAR(36),
   pair VARCHAR(10)
);

CREATE TABLE IF NOT EXISTS c2c_orders (
   id VARCHAR(36),
   order_number VARCHAR(32),
   trade_type VARCHAR(4),
   fiat VARCHAR(10),
   amount DOUBLE(10, 5),
   total_price DOUBLE(10, 5),
   create_time BITINT(8)
);

CREATE TABLE IF NOT EXISTS spot_orders (
   id VARCHAR(36),

)
`
