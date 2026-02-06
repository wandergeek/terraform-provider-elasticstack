variable "name" {
  type = string
}

provider "elasticstack" {
  kibana {}
}

resource "elasticstack_kibana_action_connector" "test_connector" {
  name           = var.name
  connector_type_id = ".index"
  config = jsonencode({
    index   = ".kibana"
    refresh = true
  })
}

resource "elasticstack_kibana_alerting_rule" "test_rule_with_actions" {
  name         = var.name
  rule_id      = "ff33ce2d-9fc4-5131-a350-b5bd6482746f"
  consumer     = "alerts"
  notify_when  = "onActiveAlert"
  rule_type_id = ".index-threshold"
  interval     = "1m"
  enabled      = true

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

  actions {
    group  = "threshold met"
    id     = elasticstack_kibana_action_connector.test_connector.connector_id
    params = jsonencode({
      documents = [
        {
          message   = "{{context.message}}"
          rule_id   = "{{rule.id}}"
          rule_name = "{{rule.name}}"
        }
      ]
    })

    frequency {
      summary    = true
      notify_when = "onActionGroupChange"
      throttle   = "10m"
    }
  }
}
