from flask import Flask, render_template, jsonify, request, url_for
app = Flask(__name__)

@app.route("/")
def home_page():
    return render_template("index.html")

@app.route("/hello")
def hello_func():
    return "Hello"

@app.route("/bal")
def bal_sheet():
    return render_template('bal_sheet.html')

@app.route("/from", methods = ['GET', 'POST'])
def form_handler():
    if request.method == 'POST':
        name = request.form.get("name")
        email = request.form.get("email")
        age = request.form.get("age")
        comments = request.form.get("comments")
        output = {
            "Aame": name,
            "Email": email,
            "Age": age,
            "comments".title(): comments
        }
        return jsonify(output)
    else:
        output2 = {"error".title(): "Method not supported"}
        jsonify(output2)

if __name__ == '__main__':
    app.run(debug=True)
