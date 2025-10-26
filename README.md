# TourismSystem
Site on vue3.js and back on Golang (gorm, chi). Non production version.
This project have back and fron ends.

## BackEnd is buided on Golang (gorm , chi).
## Fron is vue3.js (routerVue, StorePinia)

For use this sistem you need to specifay docker-compose.yaml file
fill .env with your keys.

1. Set conteiner for DB PostgresQL that work with
2. Set conteiner for BackEdn REST API on Golang (gorm, chi) that connect with DB on PostgresQL
3. FrontEndSPA on vue3.js project build it do dist directory or ather and set nginx (apache) this directory as html index