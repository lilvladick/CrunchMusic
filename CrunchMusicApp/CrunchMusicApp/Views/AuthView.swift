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
                            .stroke(Color.accentColor.opacity(0.2), lineWidth: 2)
                    )
                    .padding(.bottom, 10)
                
                SecureField("Password", text: $password)
                    .frame(height: 50)
                    .padding(.horizontal, 10)
                    .overlay(
                        RoundedRectangle(cornerRadius: 10)
                            .stroke(Color.accentColor.opacity(0.2), lineWidth: 2)
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
                .background(Color.accentColor)
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
                            .stroke(Color.accentColor.opacity(0.2), lineWidth: 2)
                    )
                    .padding(.bottom, 10)
                
                TextField("Login", text: $login)
                    .frame(height: 50)
                    .padding(.horizontal, 10)
                    .overlay(
                        RoundedRectangle(cornerRadius: 10)
                            .stroke(Color.accentColor.opacity(0.2), lineWidth: 2)
                    )
                    .padding(.bottom, 10)
                
                SecureField("Password", text: $password)
                    .frame(height: 50)
                    .padding(.horizontal, 10)
                    .overlay(
                        RoundedRectangle(cornerRadius: 10)
                            .stroke(Color.accentColor.opacity(0.2), lineWidth: 2)
                    )
                    .padding(.bottom, 20)
                NavigationLink(destination: LogInView()) {
                    HStack {
                        Text("Already have an account? ")
                            .foregroundStyle(Color.black)
                        Text("Log In")
                            .foregroundStyle(Color.accentColor)
                            .bold()
                    }.font(.callout)
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
                .background(Color.accentColor)
                .cornerRadius(10)
            }
            .padding(30)
        }
    }
}

#Preview {
    SignUpView()
}
