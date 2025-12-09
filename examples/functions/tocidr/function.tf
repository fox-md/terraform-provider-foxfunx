output "cidr" {
  value = provider::foxfunx::tocidr("10.10.10.0", "255.255.255.0")
}

# returns 10.10.10.0/24
