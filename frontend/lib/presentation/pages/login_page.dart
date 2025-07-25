import 'package:flutter/material.dart';
import 'package:inno_test/presentation/pages/registration_page.dart';
import 'package:inno_test/presentation/providers/theme_provider.dart';
import 'package:provider/provider.dart';
import '../widgets/appbars/appbar_with_text.dart';

class LoginPage extends StatefulWidget {
   const LoginPage({super.key});

  @override
  State<LoginPage> createState() => _LoginPageState();
}

class _LoginPageState extends State<LoginPage> {
  final _formKey = GlobalKey<FormState>();
  final _emailController = TextEditingController();
  final _passwordController = TextEditingController();

  bool _obscurePassword = true;

  void _submit() {
    if (_formKey.currentState!.validate()) {
      final email = _emailController.text.trim();
      final password = _passwordController.text;

      //затычка
      showDialog(
        context: context,
        builder: (_) => AlertDialog(
          title: const Text("success!"),
          content: Text("Email: $email\n password: $password"),
        ),
      );
    }
  }

  void _toSigninPage() {
     Navigator.push(
              context,
              MaterialPageRoute(builder: (context) => const RegistrationPage()),
            );
  }

  @override
  void dispose() {
    _emailController.dispose();
    _passwordController.dispose();
    super.dispose();
  }


//wait for endpoint.
  @override
  Widget build(BuildContext context) {
    final themeProvider = Provider.of<ThemeProvider>(context);
    return Scaffold(
      appBar: PreferredSize(preferredSize: Size.fromHeight(100), child: AppbarWithText(text: "Home page")),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Form(
          key: _formKey,
          child: Column(
            children: [

              LayoutBuilder(
  builder: (context, constraints) {
    double maxWidth = constraints.maxWidth > 500
        ? 500
        : constraints.maxWidth * 0.9;

    return Center(
      child: SizedBox(
        width: maxWidth,
        child: Column(
          children: [
            // Email
            Container(
              padding: const EdgeInsets.symmetric(horizontal: 15, vertical: 10),
              decoration: BoxDecoration(
                color: themeProvider.isDarkTheme
                    ? const Color(0xFF303030)
                    : const Color(0xFFF5F5F5),
                borderRadius: BorderRadius.circular(30),
                border: Border.all(
                  color: Colors.transparent,
                  width: 0.8,
                ),
              ),
              child: TextFormField(
                controller: _emailController,
                style: const TextStyle(),
                cursorColor: const Color(0xFF0088FF),
                decoration: const InputDecoration(
                  contentPadding: EdgeInsets.symmetric(vertical: 12),
                  hintStyle: TextStyle(
                    color: Color(0xFF898989),
                    fontWeight: FontWeight.w400,
                    fontFamily: "Inter",
                  ),
                  hintText: "email",
                  border: InputBorder.none,
                ),
                validator: (value) {
                  RegExp exp = RegExp(r'^\S+@\S+\.\S+$');
                  if (!exp.hasMatch(value!)) {
                    return "email is invalid";
                  }
                  return null;
                }
              ),
            ),

            const SizedBox(height: 16),

            // Password
            Container(
              padding: const EdgeInsets.symmetric(horizontal: 15, vertical: 10),
              decoration: BoxDecoration(
                color: themeProvider.isDarkTheme
                    ? const Color(0xFF303030)
                    : const Color(0xFFF5F5F5),
                borderRadius: BorderRadius.circular(30),
                border: Border.all(
                  color: Colors.transparent,
                  width: 0.8,
                ),
              ),
              child: TextFormField(
                controller: _passwordController,
                obscureText: _obscurePassword,
                style: const TextStyle(),
                cursorColor: const Color(0xFF0088FF),
                decoration: InputDecoration(
                  contentPadding: const EdgeInsets.symmetric(vertical: 12),
                  hintStyle: const TextStyle(
                    color: Color(0xFF898989),
                    fontWeight: FontWeight.w400,
                    fontFamily: "Inter",
                  ),
                  hintText: "password",
                  border: InputBorder.none,
                  suffixIcon: IconButton(
                    icon: Icon(
                      _obscurePassword
                          ? Icons.visibility
                          : Icons.visibility_off,
                    ),
                    onPressed: () {
                      setState(() {
                        _obscurePassword = !_obscurePassword;
                      });
                    },
                  ),
                ),
                validator: (value) {
                   if (value == null || value.length < 8 ) {
                    return "password should contain 8 or more characters";
                  }
                  return null;
                }
              ),
            ),
          
            const SizedBox(height: 16),

            ElevatedButton(
              onPressed: _submit,
              style: ElevatedButton.styleFrom(
                backgroundColor: Color(0xFF0088FF),
              ),
              child:
              Text("Log in",
                style: TextStyle(color: Color(0xFFF5F5F5))
              ),
            ),

             const SizedBox(height: 32),

            ElevatedButton(onPressed: _toSigninPage,
            style: ElevatedButton.styleFrom(
                backgroundColor: Color(0xFF0088FF),
              ), child: Text("Dont have account? Sigb in",
                style: TextStyle(color: Color(0xFFF5F5F5)))),
          ],
        ),
      ),
    );
  },
),
            ],
          ),
        ),
      ),
    );
  }
}