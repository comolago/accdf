python -m json.tool

curl -XDELETE 'http://localhost:9200/accdf/benchmarks/AVy32ccTJokG6S6_ipRw'
curl -XDELETE 'http://localhost:9200/accdf'


curl -i -XPOST 'http://localhost:9200/accdf/benchmarks' -d '
{
  "Id":"000000000",
  "Name":"checkconn",
  "Platforms>Platform": [
    {
          "Id":"rhel",
          "Label":"Red Hat Linux",
          "Version":"6.x"
    },
        {
          "Id":"rhel",
          "Label":"Red Hat Linux",
          "Version":"7.x"
        }
  ],
  "Requires\u003eRequire":[
    {
          "Type":"",
          "Id":"perl"
        }
  ],
  "Fingerprints":null,
  "Privileges":"\n    Cmnd_Alias                ROUTE=/sbin/route\n    Defaults!ROUTE        !requiretty\n    mcarcano  ALL=(ALL) NOPASSWD:ROUTE\n  ",
  "Description":"\n    Test connection\n  ",
  "Advice":"\n    Returns STATUS=0 if OK, STATUS=1 connection failed\n  "
}'

curl -i -XGET 'http://localhost:9200/accdf/benchmarks/_search'


curl -XGET 'http://localhost:9200/accdf/benchmarks/_search' -d '
{
  "query": {
    "bool": {
      "must": [
      { "match": { "Name": "checkconn" }},
      {
      "nested": {
        "path": "Platforms>Platform",
          "query": {
            "bool": { "must": [
              { "match": { "Platforms\u003ePlatform.Id": "rhel" }},
              { "match": { "Platforms\u003ePlatform.Version": "7.x" }}
            ]
          }
        }
      }
      }
      ]
    }
  }
}'

