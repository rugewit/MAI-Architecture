workspace {
    name "Магазин"
    description "Простая система магазина с базовыми функциями для покупателей"

    # включаем режим с иерархической системой идентификаторов
    !identifiers hierarchical

    !docs documentation
    !adrs decisions
    # Модель архитектуры
    model {

        # Настраиваем возможность создания вложенных груп
        properties { 
            structurizr.groupSeparator "/"
        }
        
        # Описание компонент модели
        user = person "Покупатель"
        shop = softwareSystem "Магазин" {
            description "Простой магазин"

            user_service = container "User service" {
                description "Сервис управления пользователями"
            }

            basket_service = container "Basket service" {
                description "Сервис управления корзиной"
            }

            product_service = container "Product service" {
                description "Сервис управления товарами"
            }

            group "Слой данных" {
                user_database = container "User Database" {
                    description "База данных с пользователями"
                    technology "PostgreSQL 15"
                    tags "database"
                }
                
                user_cache = container "User Cache" {
                    description "Кеш пользовательских данных для ускорения аутентификации"
                    technology "Redis"
                    tags "database"
                }
                
                shop_database = container "Shop Database" {
                    description "База данных для хранения информации о продуктах"
                    technology "MongoDB 5"
                    tags "database"
                }
            }

            user_service -> user_database "Получение/обновление данных о пользователях" "TCP 5432"
            user_service -> user_cache "Получение/обновление данных о пользователях" "TCP 6379"
            user_service -> basket_service  "запрос на обновление данных о корзине"
            user_service -> product_service  "запрос на обновление данных о товарах"

            product_service -> shop_database "Получение/обновление данных о товаре" "TCP 27018"
        }

        user -> shop "Покупает товар в магазине"
        user -> shop.user_service "Регистрация нового пользователя"

        deploymentEnvironment "Production" {
            deploymentNode "User Server" {
                containerInstance shop.user_service
            }

            deploymentNode "Product Server" {
                containerInstance shop.product_service
                properties {
                    "cpu" "4"
                    "ram" "256Gb"
                    "hdd" "4Tb"
                }
            }

            deploymentNode "databases" {
     
                deploymentNode "Database Server 1" {
                    containerInstance shop.user_database
                }

                deploymentNode "Database Server 2" {
                    containerInstance shop.shop_database
                    instances 3
                }

                deploymentNode "Cache Server" {
                    containerInstance shop.user_cache
                }
            }
            
        }
    }

    views {
        themes default

        properties { 
            structurizr.tooltips true
        }

        !script groovy {
            workspace.views.createDefaultViews()
            workspace.views.views.findAll { it instanceof com.structurizr.view.ModelView }.each { it.enableAutomaticLayout() }
        }

        dynamic shop "UC01" "Добавление нового пользователя" {
            autoLayout
            user -> shop.user_service "Создать нового пользователя (POST /user)"
            shop.user_service -> shop.user_database "Сохранить данные о пользователе" 
        }

        dynamic shop "UC02" "Удаление пользователя" {
            autoLayout
            user -> shop.user_service "Удалить нового пользователя (DELETE /user)"
            shop.user_service -> shop.user_database "Удалить данные о пользователе" 
        }

        dynamic shop "UC03" "Сохранить данные о товарах" {
            autoLayout
            shop.product_service -> shop.shop_database "Сохранить данные о товарах" 
        }

        dynamic shop "UC04" "Сохранить данные о корзине" {
            autoLayout
            shop.basket_service -> shop.user_service "Отправить данные о корзине"
            shop.user_service -> shop.product_service "Отправить данные о корзине"
            shop.product_service -> shop.shop_database "Сохранить данные о корзине" 
        }


        styles {
            element "database" {
                shape cylinder
            }
        }
    }
}