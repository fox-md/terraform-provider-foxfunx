output "true" {
  value = provider::foxfunx::filecontains("/etc/hosts", "localhost")
}
# returns true

output "false" {
  value = provider::foxfunx::filecontains("/etc/hosts", "hostlocal")
}
# returns false

output "case_sensitive" {
  value = provider::foxfunx::filecontains("/etc/hosts", "Localhost")
}
# returns false

output "case_insensitive" {
  value = provider::foxfunx::filecontains("/etc/hosts", "Localhost", false)
}
# returns true
