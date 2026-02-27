CREATE TABLE user_settings (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    settings JSONB
);

-- INSERT INTO user_settings(user_id , settings)
-- values(
--     1,{
--         "theme": "dark",
--         "notification:": true
--     }
--     2,{
--         "theme :" "light",
--         "notification :" false
--     }
-- )