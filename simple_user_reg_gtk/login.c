#include <stdio.h>
#include <gtk/gtk.h>
#include <string.h>

#define FILE_NAME "login_data.txt"


static void on_submit_clicked(GtkWidget *button, gpointer user_data){
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

    FILE *file = fopen(FILE_NAME, "a");
    if (file == NULL) {
        GtkWidget *dialog = gtk_message_dialog_new(
            NULL, GTK_DIALOG_MODAL, GTK_MESSAGE_ERROR, GTK_BUTTONS_OK,
            "Error: Unable to save the data!");
        gtk_dialog_run(GTK_DIALOG(dialog));
        gtk_widget_destroy(dialog);
        return;
    }

    fprintf(file, "Username: %s\nPassword: %s\n\n", username, password);
    fclose(file);

    GtkWidget *dialog = gtk_message_dialog_new(
        NULL, GTK_DIALOG_MODAL, GTK_MESSAGE_INFO, GTK_BUTTONS_OK,
        "Login information saved successfully!");
    gtk_dialog_run(GTK_DIALOG(dialog));
    gtk_widget_destroy(dialog);

    // Clear entries
    gtk_entry_set_text(GTK_ENTRY(entries[0]), "");
    gtk_entry_set_text(GTK_ENTRY(entries[1]), "");
}

int main(int argc, char *argv[]) {
    gtk_init(&argc, &argv);

    // Create the main window
    GtkWidget *window = gtk_window_new(GTK_WINDOW_TOPLEVEL);
    gtk_window_set_title(GTK_WINDOW(window), "Login");
    gtk_window_set_default_size(GTK_WINDOW(window), 300, 200);
    gtk_window_set_position(GTK_WINDOW(window), GTK_WIN_POS_CENTER);
    gtk_container_set_border_width(GTK_CONTAINER(window), 10);

    // Set a margin and vertical box
    GtkWidget *vbox = gtk_box_new(GTK_ORIENTATION_VERTICAL, 10);
    gtk_container_add(GTK_CONTAINER(window), vbox);

    // Add a welcome label
    GtkWidget *label = gtk_label_new("Welcome! Please Log In");
    gtk_box_pack_start(GTK_BOX(vbox), label, FALSE, FALSE, 0);

    // Username entry
    GtkWidget *username_entry = gtk_entry_new();
    gtk_entry_set_placeholder_text(GTK_ENTRY(username_entry), "Username");
    gtk_box_pack_start(GTK_BOX(vbox), username_entry, FALSE, FALSE, 0);

    // Password entry
    GtkWidget *password_entry = gtk_entry_new();
    gtk_entry_set_placeholder_text(GTK_ENTRY(password_entry), "Password");
    gtk_entry_set_visibility(GTK_ENTRY(password_entry), FALSE); // Hide password
    gtk_box_pack_start(GTK_BOX(vbox), password_entry, FALSE, FALSE, 0);

    // Submit button
    GtkWidget *submit_button = gtk_button_new_with_label("Submit");
    gtk_box_pack_start(GTK_BOX(vbox), submit_button, FALSE, FALSE, 0);

    // Array to store the entries
    GtkWidget *entries[2] = {username_entry, password_entry};

    // Connect the button signal to callback
    g_signal_connect(submit_button, "clicked", G_CALLBACK(on_submit_clicked), entries);

    // Connect the window destroy signal
    g_signal_connect(window, "destroy", G_CALLBACK(gtk_main_quit), NULL);

    // Show all widgets
    gtk_widget_show_all(window);

    // Run the main GTK loop
    gtk_main();

    return 0;
}