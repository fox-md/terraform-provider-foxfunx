# Check if /etc directory exists
output "etc" {
  value = provider::foxfunx::direxists("/etc")
}

# Check if /etc/hosts directory exists
output "hosts" {
  value = provider::foxfunx::direxists("/etc/hosts")
}
