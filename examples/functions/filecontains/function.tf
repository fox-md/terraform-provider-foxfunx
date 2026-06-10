output "true" {
  value = provider::foxfunx::filecontains("/etc/hosts", "localhost")
}
# returns true

output "false" {
  value = provider::foxfunx::filecontains("/etc/hosts", "hostlocal")
}
# returns false
