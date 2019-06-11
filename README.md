# es

mkdir -p ~/go/src/github.com/es/

cd ~/go/src/github.com/es/
git clone git@github.com:TretyakovArtem/es.git .


для создания бд необходимо запустить в терминале команду из корня проекта

`mysql -u username -p < db.sql`

get запрос по `/` сгенерирует 10 000 записей в трех таблицах `oms_orders`

настройка подключения к бд находится в `driver/driver.go`
