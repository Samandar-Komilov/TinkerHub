#include <stdio.h>
#include <stdlib.h>

typedef struct Node {
    int data;
    struct Node *left;
    struct Node *right;
} Node;

Node* createNode(int data){
    Node* newNode = (Node*) malloc(sizeof(Node));
    newNode->data = data;
    newNode->right = NULL;
    newNode->left = NULL;

    return newNode;
}

Node* insert(Node* root, int data){
    if (root == NULL) return createNode(data);

    if (data < root->data)
        root->left = insert(root->left, data);
    else
        root->right = insert(root->right, data);

    return root;
}

Node* search(Node* root, int key){
    if (root == NULL || root->data == key) return root;

    if (key > root->data) return search(root->right, key);

    return search(root->left, key);
}

void inOrderTraversal(Node* root){
    if (root != NULL){
        inOrderTraversal(root->left);
        printf("%d ", root->data);
        inOrderTraversal(root->right);
    }
}




int main()
{
    // Soon, BST implementation will be written
    Node* root = NULL;
    root = insert(root, 17);
    insert(root, 30);
    insert(root, 20);

    printf("Inorder traversal of the BST: ");
    inOrderTraversal(root);

    return 0;
}
