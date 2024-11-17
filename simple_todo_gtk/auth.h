#ifndef AUTH_H

#include <gtk/gtk.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int user_exists(const char *username);
int validate_login(const char *username, const char *password);
void save_user(const char *username, const char *password);
void on_register_clicked(GtkWidget *button, gpointer user_data);
void on_login_clicked(GtkWidget *button, gpointer user_data);

#endif