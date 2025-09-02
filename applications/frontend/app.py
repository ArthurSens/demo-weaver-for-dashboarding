from flask import Flask, render_template, redirect, flash, g, request, Response
from werkzeug.middleware.dispatcher import DispatcherMiddleware
from prometheus_client import make_wsgi_app
from o11y_lib import http
import requests
import time

app = Flask(__name__)
app.secret_key = b'a-very-good-secret'
app.wsgi_app = DispatcherMiddleware(app.wsgi_app, {
    '/metrics': make_wsgi_app()
})

durations = http.ServerRequestDuration()

@app.before_request
def start_timer():
    g.method = request.method
    g.scheme = request.scheme
    g.route = request.path
    g.start = time.time()

@app.after_request
def log_details(response: Response):
    durations.with_attr(g.method,g.scheme,response_status_code=response.status,route=g.route).observe(time.time() - g.start)
    return response

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
