import SwiftUI

struct LogInView: View {
    @State private var email: String = ""
    @State private var password: String = ""
    
    var body: some View {
        NavigationStack {
            VStack {
                Spacer()
                
                Text("Log In")
                    .font(.largeTitle)
                    .bold()
                    .padding(.bottom, 20)
                
                TextField("Email", text: $email)
                    .frame(height: 50)
                    .padding(.horizontal, 10)
                    .overlay(
                        RoundedRectangle(cornerRadius: 10)
                            .stroke(Color.gray.opacity(0.2), lineWidth: 2)
                    )
                    .padding(.bottom, 10)
                
                SecureField("Password", text: $password)
                    .frame(height: 50)
                    .padding(.horizontal, 10)
                    .overlay(
                        RoundedRectangle(cornerRadius: 10)
                            .stroke(Color.gray.opacity(0.2), lineWidth: 2)
                    )
                    .padding(.bottom, 20)
                
                Spacer()
                
                Button(action: {
                    // login func
                }, label: {
                    Text("Log In")
                        .bold()
                        .foregroundColor(.white)
                        .padding(10)
                        .frame(maxWidth: 100)
                })
                .background(Color.blue)
                .cornerRadius(10)
            }
            .padding(30)
        }
    }
}

struct SignUpView: View {
    @State private var email: String = ""
    @State private var password: String = ""
    @State private var login: String = ""
    
    var body: some View {
        NavigationStack {
            VStack {
                Spacer()
                
                Text("Sign Up")
                    .font(.largeTitle)
                    .bold()
                    .padding(.bottom, 20)
                
                TextField("Email", text: $email)
                    .frame(height: 50)
                    .padding(.horizontal, 10)
                    .overlay(
                        RoundedRectangle(cornerRadius: 10)
                            .stroke(Color.gray.opacity(0.2), lineWidth: 2)
                    )
                    .padding(.bottom, 10)
                
                TextField("Login", text: $login)
                    .frame(height: 50)
                    .padding(.horizontal, 10)
                    .overlay(
                        RoundedRectangle(cornerRadius: 10)
                            .stroke(Color.gray.opacity(0.2), lineWidth: 2)
                    )
                    .padding(.bottom, 10)
                
                SecureField("Password", text: $password)
                    .frame(height: 50)
                    .padding(.horizontal, 10)
                    .overlay(
                        RoundedRectangle(cornerRadius: 10)
                            .stroke(Color.gray.opacity(0.2), lineWidth: 2)
                    )
                    .padding(.bottom, 20)
                NavigationLink(destination: LogInView()) {
                    Text("Already have an account? Log In")
                        .font(.callout)
                        .foregroundStyle(Color.blue)
                }
                
                Spacer()
                
                Button(action: {
                    // save account to keychain database + save to server
                }, label: {
                    Text("Sign Up")
                        .bold()
                        .foregroundColor(.white)
                        .padding(10)
                        .frame(maxWidth: 100)
                })
                .background(Color.blue)
                .cornerRadius(10)
            }
            .padding(30)
        }
    }
}

#Preview {
    SignUpView()
}
