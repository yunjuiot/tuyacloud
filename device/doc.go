// Package device implements Device Management API interface.
package device

// Implemented:
// GET		/v1.0/devices/{device_id}							Get device details
// GET		/v1.0/users/{uid}/devices							Get the list of devices under the user
// GET		/v1.0/devices										Get device list
// PUT		/v1.0/devices/{device_id}/functions/{function_code}	Modify the function point name
// GET		/v1.0/devices/{device_id}/logs						Query device logs
// PUT		/v1.0/devices/{device_id}/reset-factory				Restore the factory settings
// DELETE	/v1.0/devices/{device_id}							Delete device
// GET		/v1.0/devices/{deviceId}/sub-devices				Query the list of devices under the gateway
// GET		/v1.0/devices/factory-infos							Query device factory information