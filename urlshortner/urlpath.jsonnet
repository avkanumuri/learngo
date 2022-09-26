local services = ["ingress", "thanos", "prom", "alerts", "jaeger"];
local edges= ["use1a", "use1b", "use1c", "usw2a", "usw2b"];
local edgeurl = [ {
    path: "/" + edge_name + "_" + service_name, 
    url: "https://" + service_name + '.' + edge_name+ ".aws.parsec.apple.com"
}
    for edge_name in edges
    for service_name in services
];

{
    "Path2URL.yaml" : std.manifestYamlDoc(
  edgeurl
)
}
