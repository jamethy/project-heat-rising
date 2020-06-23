variable database_url {
  type        = string
  description = "Full URL of database"
}

variable database_username {
  type        = string
  description = "App username of database"
}

variable database_password {
  type        = string
  description = "App password of database"
}

variable carrier_username {
  type        = string
  description = "Username for carrier COR thermostat."
  default     = ""
}

variable carrier_password {
  type        = string
  description = "Password for carrier COR thermostat"
  default     = ""
}

variable open_weather_base_url {
  type        = string
  description = "OpenWeatherMaps base url"
  default     = ""
}

variable open_weather_api_key {
  type        = string
  description = "OpenWeatherMaps API Key"
  default     = ""
}

variable open_weather_lat {
  type        = string
  description = "OpenWeatherMaps desired weather location latitude"
  default     = ""
}

variable open_weather_lon {
  type        = string
  description = "OpenWeatherMaps desired weather location longitude"
  default     = ""
}
