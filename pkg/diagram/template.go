package diagram

var begin string = `
@startuml C4_Elements
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml
`

var end string = `@enduml`

var label string = `%s(%s, "%s", "%s")`
var rel string = `Rel(%s, %s, "%s", "%s")`
