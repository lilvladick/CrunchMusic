import SwiftUI

struct SettingsView: View {
    var body: some View {
        NavigationStack {
            Form {
                Button("Delete Uploadet Tracks") {
                    
                }
                Button("Change Password") {
                    
                }
                Button("Delete Account") {
                    
                }.foregroundStyle(Color.red)
            }
            .navigationTitle("Settings")
        }
    }
}


#Preview {
    SettingsView()
}
