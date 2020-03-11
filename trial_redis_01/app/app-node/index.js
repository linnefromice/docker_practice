const Redis = require('ioredis');
const redis = new Redis();

const main = async () => {
    await redis.set('KEY', 'set_value');
    const result = await redis.get('KEY');
    console.log(result);
    redis.disconnect();
}

main();