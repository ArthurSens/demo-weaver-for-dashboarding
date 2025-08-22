from flask import Flask, render_template, redirect, url_for, flash
import requests

app = Flask(__name__)
app.secret_key = b'a-very-good-secret'

def get_books():
    response = requests.get("http://localhost:8080/books")
    return response.json()

@app.route("/")
def list_books():
    books = get_books()
    return render_template('list.html',books=books)

@app.route("/books/<book_id>/borrow",methods=["POST"])
def borrow_book(book_id):
    requests.post(f"http://localhost:8080/books/{book_id}/borrow")
    flash("Book borrowed!","info")
    return redirect('/')

@app.route("/books/<book_id>/return",methods=["POST"])
def return_book(book_id):
    requests.post(f"http://localhost:8080/books/{book_id}/return")
    flash("Book returned!","info")
    return redirect('/')
