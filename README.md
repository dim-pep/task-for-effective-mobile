# Тестовое задание Effective Mobile

REST‑сервис для агрегации данных об онлайн‑подписках пользователей.

Реализовано:  
- CRUDL‑операции над подписками  
- Ручка для подсчета суммарной стоимости за выбранный период с фильтрами по пользователю и названию сервиса.
- Миграции, логирование, конфигурация через .env, Swagger и запуск через Docker Compose.

## Как запустить
1) Создайте .env в корне проекта и скопируй значения из .env.example, при желании можно отредактировать.
2) Запусти docker:
```
docker compose up -d --build
```
3) Логи можно посмотреть внутри контейнера subscriptions-app или командой:
```
docker compose logs -f app
```
4) Swagger доступен по ссылке:  
http://localhost:8080/swagger/index.html#/

5) Если вы не меняли значения .env.example, то подключиться к бд можно внтури контейнера subscriptions-db, командой: 
```
psql -U postgres -d subscriptions_db
```

## API
POST /subscriptions — создать подписку  
GET /subscriptions/list — список uid подписок  
GET /subscriptions/{id} — получить подписку по UID  
PUT /subscriptions/{id} — полностью обновить подписку по UID  
DELETE /subscriptions/{id} — удалить подписку  по UID  
GET /subscriptions/summary — суммарная стоимость подписок за период  

Через Swagger доступны все эндоинты
