endpoint "typetest" "http" {
  endpoint = "/"
  method   = "POST"
  codec    = "json"
}

endpoint "typetest" "grpc" {
  package = "semaphore.typetest"
  service = "Typetest"
  method  = "Run"
}

flow "typetest" {
  input "semaphore.typetest.Request" {}

  output "semaphore.typetest.Response" {
    data = {
      empty = "{{ input:data.empty }}"
      single = "{{ input:data.single }}"
      plural = "{{ input:data.plural }}"
    }

    nested = "{{ input:nested }}"
  }
}