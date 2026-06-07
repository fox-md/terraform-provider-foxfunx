# Check if /etc directory is empty
output "etc" {
  value = provider::foxfunx::dirempty("/etc")
}

# false

# Check if /etc/hosts directory is empty
output "hosts" {
  value = provider::foxfunx::dirempty("/etc/hosts")
}

# returns an error as `/etc/hosts` is a file
