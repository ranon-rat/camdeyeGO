#include <Arduino.h>
#include <WiFi.h>
#include "network_config.hpp"
void setup()
{
  Serial.begin(115200);
  Serial.setDebugOutput(false);
  // Set pin mode
  WiFi.config(INADDR_NONE, INADDR_NONE, INADDR_NONE, INADDR_NONE);
  WiFi.mode(WIFI_STA);
  WiFi.setHostname("camdeyeCam");

  WiFi.begin((char *)SSID, (char *)PASS); // connect to our own network

  while (WiFi.status() != WL_CONNECTED)
  {
    delay(500);
    Serial.print(".");
  }

  Serial.printf("\nWiFi connected\n\nCamera Stream Ready! Go to: http://%s/\n", WiFi.localIP().toString().c_str());
}
void loop()
{

  // put your main code here, to run repeatedly:
}