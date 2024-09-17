import SwiftUI

struct SettingsView: View {
    //@AppStorage private var darkMode: Bool = false
    var body: some View {
        NavigationStack {
            Form {
                //Toggle("Dark Mode", isOn: $darkMode)
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
