package main

import (
	"fmt"
	"github.com/danielstutzman/go-monitis"
	"github.com/hashicorp/terraform/helper/schema"
	"strconv"
)

func resource_monitis_external_monitor() *schema.Resource {
	return &schema.Resource{
		Create: resource_monitis_external_monitor_create,
		Read:   resource_monitis_external_monitor_read,
		Update: resource_monitis_external_monitor_update,
		Delete: resource_monitis_external_monitor_delete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func convertDataToOptions(d *schema.ResourceData) *monitis.AddExternalMonitorOptions {
	opts := monitis.AddExternalMonitorOptions{}

	if name, ok := d.Get("name").(string); ok {
		opts.Name = monitis.String(name)
	}

	return &opts
}

func resource_monitis_external_monitor_create(d *schema.ResourceData, m interface{}) error {
	auth := m.(*monitis.Auth)
	opts := convertDataToOptions(d)

	monitorData, err := auth.AddExternalMonitor(opts)
	if err != nil {
		return fmt.Errorf("Error from AddExternalMonitor: %s", err)
	}

	d.SetId(strconv.Itoa(monitorData.TestId))

	return err
}

func resource_monitis_external_monitor_read(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resource_monitis_external_monitor_update(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resource_monitis_external_monitor_delete(d *schema.ResourceData, m interface{}) error {
	auth := m.(*monitis.Auth)
	testId := d.Id()

	err := auth.DeleteExternalMonitors(testId, nil)

	return err
}
