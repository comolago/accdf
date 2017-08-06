
Get Endpopints IP

xpath topology.xml "/Topology/Endpoints/Endpoint[@id=/Topology/Interconnections/Interconnection[@systemid=/Topology/Systems/System[@name='myhost01']/@id]/@endpointid]"/@ip

Get Endpoints Ports

xpath topology.xml "/Topology/Endpoints/Endpoint[@id=/Topology/Interconnections/Interconnection[@systemid=/Topology/Systems/System[@name='myhost01']/@id]/@endpointid]"/@port

Get Protocols

xpath topology.xml "/Topology/Services/Service[@id=/Topology/Bindings/Binding[@endpointid=/Topology/Endpoints/Endpoint[@id=/Topology/Interconnections/Interconnection[@systemid=/Topology/Systems/System[@name='myhost01']/@id]/@endpointid]/@id]/@serviceid]/@protocol"

Get Entities
xpath topology.xml "/Topology/Entities/Entity[@id=/Topology/ServiceDetails/ServiceDetail[@serviceid=/Topology/Services/Service[@id=/Topology/Bindings/Binding[@endpointid=/Topology/Endpoints/Endpoint[@id=/Topology/Interconnections/Interconnection[@systemid=/Topology/Systems/System[@name='myhost01']/@id]/@endpointid]/@id]/@serviceid]/@id]/@entityid]"
