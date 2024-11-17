#include "auth.h"

#define USERS_FILE "users.txt"

// Function to check if a username exists in the file
int user_exists(const char *username) {
    FILE *file = fopen(USERS_FILE, "r");
    if (!file) return 0;

    char line[256];
    while (fgets(line, sizeof(line), file)) {
        char stored_username[128];
        sscanf(line, "%127[^:]", stored_username);
        if (strcmp(username, stored_username) == 0) {
            fclose(file);
            return 1;
        }
    }
    fclose(file);
    return 0;
}

// Function to validate login credentials
int validate_login(const char *username, const char *password) {
    FILE *file = fopen(USERS_FILE, "r");
    if (!file) return 0;

    char line[256];
    while (fgets(line, sizeof(line), file)) {
        char stored_username[128], stored_password[128];
        sscanf(line, "%127[^:]:%127s", stored_username, stored_password);
        if (strcmp(username, stored_username) == 0 && strcmp(password, stored_password) == 0) {
            fclose(file);
            return 1;
        }
    }
    fclose(file);
    return 0;
}

// Function to save a new user to the file
void save_user(const char *username, const char *password) {
    FILE *file = fopen(USERS_FILE, "a");
    if (file) {
        fprintf(file, "%s:%s\n", username, password);
        fclose(file);
    }
}

// Callback for register button
void on_register_clicked(GtkWidget *button, gpointer user_data) {
    GtkWidget **entries = (GtkWidget **)user_data;
    const char *username = gtk_entry_get_text(GTK_ENTRY(entries[0]));
    const char *password = gtk_entry_get_text(GTK_ENTRY(entries[1]));

    if (strlen(username) == 0 || strlen(password) == 0) {
        GtkWidget *dialog = gtk_message_dialog_new(
            NULL, GTK_DIALOG_MODAL, GTK_MESSAGE_WARNING, GTK_BUTTONS_OK,
            "Username and Password cannot be empty!");
        gtk_dialog_run(GTK_DIALOG(dialog));
        gtk_widget_destroy(dialog);
        return;
    }

    if (user_exists(username)) {
        GtkWidget *dialog = gtk_message_dialog_new(
            NULL, GTK_DIALOG_MODAL, GTK_MESSAGE_ERROR, GTK_BUTTONS_OK,
            "User already exists! Please choose a different username.");
        gtk_dialog_run(GTK_DIALOG(dialog));
        gtk_widget_destroy(dialog);
        return;
    }

    save_user(username, password);

    GtkWidget *dialog = gtk_message_dialog_new(
        NULL, GTK_DIALOG_MODAL, GTK_MESSAGE_INFO, GTK_BUTTONS_OK,
        "Registration successful!");
    gtk_dialog_run(GTK_DIALOG(dialog));
    gtk_widget_destroy(dialog);

    gtk_entry_set_text(GTK_ENTRY(entries[0]), "");
    gtk_entry_set_text(GTK_ENTRY(entries[1]), "");
}

// Callback for login button
void on_login_clicked(GtkWidget *button, gpointer user_data) {
    GtkWidget **entries = (GtkWidget **)user_data;
    const char *username = gtk_entry_get_text(GTK_ENTRY(entries[0]));
    const char *password = gtk_entry_get_text(GTK_ENTRY(entries[1]));

    if (validate_login(username, password)) {
        GtkWidget *dialog = gtk_message_dialog_new(
            NULL, GTK_DIALOG_MODAL, GTK_MESSAGE_INFO, GTK_BUTTONS_OK,
            "Login successful!");
        gtk_dialog_run(GTK_DIALOG(dialog));
        gtk_widget_destroy(dialog);

        // Proceed to the main to-do list screen here
        gtk_main_quit(); // Placeholder: Quit for now
    } else {
        GtkWidget *dialog = gtk_message_dialog_new(
            NULL, GTK_DIALOG_MODAL, GTK_MESSAGE_ERROR, GTK_BUTTONS_OK,
            "Invalid username or password!");
        gtk_dialog_run(GTK_DIALOG(dialog));
        gtk_widget_destroy(dialog);
    }
}
