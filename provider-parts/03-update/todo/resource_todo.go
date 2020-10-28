package todo

import (
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/myuser/todo-terraform-provider-class/client"
	"github.com/myuser/todo-terraform-provider-class/client/todos"
	"github.com/myuser/todo-terraform-provider-class/models"
)

// resourceTodo returns the resourceTodo resource schema
func resourceTodo() *schema.Resource {
	return &schema.Resource{
		Read:   resourceTodoRead,
		Update: resourceTodoUpdate,

		Schema: map[string]*schema.Schema{
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"completed": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
		},
	}
}

//resourceTodoRead reads a single todo resource and updates the state
func resourceTodoRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.TodoList)

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		log.Printf("[DEBUG] %s", err)
		return err
	}

	params := todos.NewFindTodoParams()
	params.SetID(int64(id))
	result, err := c.Todos.FindTodo(params)

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		d.SetId("")
		return nil
	}

	item := result.GetPayload()
	description := item[0].Description
	completed := item[0].Completed

	// Tell terraform what we got back from the upstream API
	err = d.Set("description", &description)
	if err != nil {
		return err
	}
	err = d.Set("completed", &completed)
	if err != nil {
		return err
	}
	return nil
}

//resourceTodoUpdate modifies a single todo resource and updates the state
func resourceTodoUpdate(d *schema.ResourceData, m interface{}) error {
	// Enable partial state mode
	// We don't really need this since we can update
	// everything with a single call.
	//d.Partial(true)

	c := m.(*client.TodoList)

	if d.HasChange("description") || d.HasChange("completed") {
		// Try updating the todo
		description := d.Get("description").(string)
		completed := d.Get("completed").(bool)

		id, err := strconv.Atoi(d.Id())
		if err != nil {
			log.Printf("[DEBUG] %s", err)
			return err
		}

		var todo = models.Item{Completed: &completed, Description: &description}
		params := todos.NewUpdateOneParams()
		params.SetBody(&todo)
		params.SetID(int64(id))
		_, err = c.Todos.UpdateOne(params)
		if err != nil {
			log.Printf("[DEBUG] %s", err)
			return err
		}

		// If we had to handle description and completion with different
		// API calls and we were to return only after updating description,
		// and before disabling partial mode below,
		// then only the "description" field would be saved.

		//d.SetPartial("description")
	}

	// We succeeded, disable partial mode.
	// This causes Terraform to save all fields again.
	//d.Partial(false)

	return resourceTodoRead(d, m)
}
