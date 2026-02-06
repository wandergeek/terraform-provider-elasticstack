variable "name" {
  type = string
}

provider "elasticstack" {
  kibana {}
}

resource "elasticstack_kibana_alerting_rule" "test_rule_disabled" {
  name         = var.name
  rule_id      = "gg33ce2d-9fc4-5131-a350-b5bd6482747g"
  consumer     = "alerts"
  notify_when  = "onActiveAlert"
  rule_type_id = ".index-threshold"
  interval     = "1m"
  enabled      = false

  params = jsonencode({
    "index" : [".test-index"],
    "timeField" : "@timestamp",
    "aggType" : "count",
    "groupBy" : "all",
    "timeWindowSize" : 5,
    "timeWindowUnit" : "m",
    "thresholdComparator" : ">",
    "threshold" : [1000]
  })
}
