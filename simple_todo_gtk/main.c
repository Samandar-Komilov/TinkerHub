#include <gtk/gtk.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#include "auth.h"

#define USERS_FILE "users.txt"


int main(int argc, char *argv[]) {
    gtk_init(&argc, &argv);

    // Main window
    GtkWidget *window = gtk_window_new(GTK_WINDOW_TOPLEVEL);
    gtk_window_set_title(GTK_WINDOW(window), "User Authentication");
    gtk_window_set_default_size(GTK_WINDOW(window), 400, 300);
    gtk_window_set_position(GTK_WINDOW(window), GTK_WIN_POS_CENTER);
    gtk_container_set_border_width(GTK_CONTAINER(window), 10);

    GtkWidget *notebook = gtk_notebook_new();
    gtk_container_add(GTK_CONTAINER(window), notebook);

    // Register Page
    GtkWidget *register_page = gtk_box_new(GTK_ORIENTATION_VERTICAL, 10);
    GtkWidget *register_username = gtk_entry_new();
    gtk_entry_set_placeholder_text(GTK_ENTRY(register_username), "Username");
    GtkWidget *register_password = gtk_entry_new();
    gtk_entry_set_placeholder_text(GTK_ENTRY(register_password), "Password");
    gtk_entry_set_visibility(GTK_ENTRY(register_password), FALSE);
    GtkWidget *register_button = gtk_button_new_with_label("Register");
    GtkWidget *register_entries[] = {register_username, register_password};
    g_signal_connect(register_button, "clicked", G_CALLBACK(on_register_clicked), register_entries);

    gtk_box_pack_start(GTK_BOX(register_page), register_username, FALSE, FALSE, 0);
    gtk_box_pack_start(GTK_BOX(register_page), register_password, FALSE, FALSE, 0);
    gtk_box_pack_start(GTK_BOX(register_page), register_button, FALSE, FALSE, 0);

    // Login Page
    GtkWidget *login_page = gtk_box_new(GTK_ORIENTATION_VERTICAL, 10);
    GtkWidget *login_username = gtk_entry_new();
    gtk_entry_set_placeholder_text(GTK_ENTRY(login_username), "Username");
    GtkWidget *login_password = gtk_entry_new();
    gtk_entry_set_placeholder_text(GTK_ENTRY(login_password), "Password");
    gtk_entry_set_visibility(GTK_ENTRY(login_password), FALSE);
    GtkWidget *login_button = gtk_button_new_with_label("Login");
    GtkWidget *login_entries[] = {login_username, login_password};
    g_signal_connect(login_button, "clicked", G_CALLBACK(on_login_clicked), login_entries);

    gtk_box_pack_start(GTK_BOX(login_page), login_username, FALSE, FALSE, 0);
    gtk_box_pack_start(GTK_BOX(login_page), login_password, FALSE, FALSE, 0);
    gtk_box_pack_start(GTK_BOX(login_page), login_button, FALSE, FALSE, 0);

    // Add pages to the notebook
    gtk_notebook_append_page(GTK_NOTEBOOK(notebook), register_page, gtk_label_new("Register"));
    gtk_notebook_append_page(GTK_NOTEBOOK(notebook), login_page, gtk_label_new("Login"));

    g_signal_connect(window, "destroy", G_CALLBACK(gtk_main_quit), NULL);

    gtk_widget_show_all(window);
    gtk_main();

    return 0;
}
