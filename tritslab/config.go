package tritslab

const TD = false // Trits debug on/off

const NOMINAL = 1           // Nominal, if set to 0 then nominal is random 1-10 % of the pocket
const GAMES_ON_TABLE = 23   // Max 23
const PLAYERS_IN_SQUAD = 14 // Max 14
const BONUS_LOW = 2
const BONUS_HIGH = 4
const ZION_RECHARGE_AT = 0.9 // Bank helps out zion players who go under this treshold
const LOG_LEVEL = LOG_DEBUG  // LOG_DEBUG<LOG_INFO < LOG_NOT
const LOG_FILE = "/var/log/trits/tritslab.log"
const WORLDS_MONEY = 1000000000 // Borrow fund initial amount
const PROFIT_TRESHOLD = 4       // Times the initial amount
const LEAVE_GAME_PROB = 50      // 1-100 chance to leave the game if profit is gretter than the PROFIT_TRESHOLD

// Influx related stuff
const INFLUX = true                                                                                               // Write statistics to influx
const INFLUX_HOST = "127.0.0.1"                                                                                   // Influx DB host
const INFLUX_PORT = "8086"                                                                                        // Influx DB port
const INFLUX_API_KEY = "BwRDdZLL2kRaRJvfzjrb-Yw-cOBph2R7JLt5YWS-6BkZp5s8BKucn_bZ4vskHJNI29vrldyGwPn6oU-z2mKfiQ==" // Influx DB API KEY
const INFLUX_ORG = "Trits"                                                                                        // Influx DB Org
const INFLUX_BUCKET = "tritslab"                                                                                  // Influx DB bucket
const INFLUX_FLUSH_EVERY = 10000                                                                                  // Flush to influx every xth write
