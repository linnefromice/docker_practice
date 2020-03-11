const MongoClient = require('mongodb').MongoClient;
const assert = require('assert');

const url = 'mongodb://localhost:27017';
const dbName = 'myMongo';

const connectOption = {
    useNewUrlParser: true,
    useUnifiedTopology: true,
}

MongoClient.connect(url, connectOption, (err, client) => {

  assert.equal(null, err);

  console.log('Connected successfully to server');

  const db = client.db(dbName);
  client.close();
});