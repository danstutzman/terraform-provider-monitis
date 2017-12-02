package main

import (
	"fmt"
	"github.com/danielstutzman/go-monitis"
	"github.com/hashicorp/terraform/helper/schema"
	"strconv"
)

func data_source_monitis_external_location() *schema.Resource {
	return &schema.Resource{
		Read: data_source_monitis_external_location_read,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"full_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"host_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"min_check_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"group": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"city": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func data_source_monitis_external_location_read(d *schema.ResourceData,
	m interface{}) error {
	auth := m.(*monitis.Auth)
	allLocations, err := auth.GetLocations()

	if err != nil {
		return fmt.Errorf("Error from GetLocations: %s", err)
	}

	nameFilter, hasNameFilter := d.GetOk("name")
	if !hasNameFilter {
		return fmt.Errorf("You must specify a name to filter by")
	}

	filteredLocations := []monitis.ExternalLocation{}
	for _, location := range allLocations {
		if location.Name == nameFilter {
			filteredLocations = append(filteredLocations, location)
		}
	}

	if len(filteredLocations) == 0 {
		return fmt.Errorf("no matching locations found")
	}
	if len(filteredLocations) > 1 {
		return fmt.Errorf("multiple regions matched; use additional constraints to reduce matches to a single region")
	}
	location := filteredLocations[0]

	d.SetId(strconv.Itoa(location.Id))
	d.Set("name", location.Name)

	return nil
}
