# Directory and Package Naming Convention
app/
    database.go
controller/
    buyer/ [buyer_controller]
    event/ [event_controller]
    seller/ [seller_controller]
db/
    connection/
    migration/
helper/
model/
    domain/
    entity/
        buyer/ [buyer_entity]
        event/ [event_entity]
        seller/ [seller_entity]
    web/
        buyer/ [buyer_web]
        event/ [event_web]
        seller/ [seller_web]
query_builder/
    buyer/ [buyer_query_builder]
    seller/
repository/
    buyer/ [buyer_repository]
    event/ [event_repository]
    seller/ [seller_repository]
route/
service/
    buyer/ [buyer_service]
    event/ [event_service]
    seller/ [seller_service]


# Type Naming Convention
interface -> implementation
BuyerController -> BuyerControllerImpl


# Extending Type Convention
func (controller *UserControllerImpl) Method()
