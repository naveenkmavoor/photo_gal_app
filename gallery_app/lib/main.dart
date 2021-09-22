import 'package:flutter/material.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: const MyHomePage(),
    );
  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({Key? key}) : super(key: key);

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  Map<String, dynamic> credentials = {};
  String err = '';
  final _formKey = GlobalKey<FormState>();
  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          title: const Text(
            'Gallery App',
            textAlign: TextAlign.center,
          ),
        ),
        body: SingleChildScrollView(
          child: Container(
            alignment: Alignment.center,
            margin: const EdgeInsets.all(10.0),
            child: Form(
              key: _formKey,
              child: Column(
                children: [
                  _buildEmailForm(),
                  const SizedBox(height: 10.0),
                  _buildPasswordForm(),
                  const SizedBox(height: 10.0),
                  _buildSubmitForm()
                ],
              ),
            ),
          ),
        ));
  }

  Widget _buildEmailForm() {
    return TextFormField(
      onSaved: (val) {
        if (val == null || val.isEmpty) {
          err = 'Please enter the email address';
        } else if (!RegExp(
                r"^[a-zA-Z0-9.a-zA-Z0-9.!#$%&'*+-/=?^_`{|}~]+@[a-zA-Z0-9]+\.[a-zA-Z]+")
            .hasMatch(val)) {
          err = 'Invalid email address';
        } else {
          credentials['email'] = val;
        }
      },
      keyboardType: TextInputType.emailAddress,
      decoration: InputDecoration(
          hintText: 'Email address',
          border: InputBorder.none,
          focusedBorder: OutlineInputBorder(
              borderRadius: BorderRadius.circular(25.0),
              borderSide: const BorderSide(color: Colors.red)),
          enabledBorder: OutlineInputBorder(
              borderRadius: BorderRadius.circular(25.0),
              borderSide: const BorderSide(color: Colors.blue, width: 2.0))),
    );
  }

  Widget _buildPasswordForm() {
    return TextFormField(
      obscureText: true,
      onSaved: (val) {
        if (val == null || val.isEmpty || val.length < 6) {
          err = 'Invalid password';
        } else {
          credentials['password'] = val;
        }
      },
      decoration: InputDecoration(
          hintText: 'Password',
          border: InputBorder.none,
          focusedBorder: OutlineInputBorder(
              borderRadius: BorderRadius.circular(25.0),
              borderSide: const BorderSide(color: Colors.red)),
          enabledBorder: OutlineInputBorder(
              borderRadius: BorderRadius.circular(25.0),
              borderSide: const BorderSide(color: Colors.blue, width: 2.0))),
    );
  }

  Widget _buildSubmitForm() {
    return ElevatedButton(
      onPressed: () {
        _formKey.currentState!.save();
        if (err == "") {
          ScaffoldMessenger.of(context).showSnackBar(SnackBar(
              content: Text(
                  'email : ${credentials['email']}, password : ${credentials['password']}')));
        } else {
          ScaffoldMessenger.of(context).showSnackBar(
              const SnackBar(content: Text('Oops! something goes wrong.')));
        }
        err = "";
      },
      child: const Text("Submit"),
    );
  }
}
