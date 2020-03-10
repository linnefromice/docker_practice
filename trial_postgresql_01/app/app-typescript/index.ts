import { Client } from 'pg';

type DatabaseSetting = {
  host: string;
  port: number;
  user: string;
  password: string;
  database: string;
};

class Postgres {
  static async getClient(setting: DatabaseSetting) {
    const { host, port, user, password, database } = setting;
    const client = new Client({
      host,
      port,
      user,
      password,
      database
    });
    await client.connect();
    return client;
  }
}

async function pgSampleQuery(client: Client) {
  return client.query('select * from users');
}

async function main() {
  const pgClient = await Postgres.getClient({
    host: 'localhost',
    port: 15432,
    user: 'example1',
    password: 'example1',
    database: 'example1'
  });

  const result = await pgSampleQuery(pgClient);
  console.log(result.rows);
}

// execute
main();