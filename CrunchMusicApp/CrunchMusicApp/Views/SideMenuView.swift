import SwiftUI

struct SideMenuView: View {

    var body: some View {
        VStack(alignment: .leading) {
            // load track
            Button(action: {
                
            }, label: {
                Image(systemName: "square.and.arrow.up")
                Text("Upload Tracks")
            }).padding(.vertical, 15)
            
            // my tracks (in view can pick your tracks only or not)
            Button(action: {
                
            }, label: {
                Image(systemName: "music.note.list")
                Text("Tracks")
            }).padding(.bottom, 15)
            
            // Settings
            Button(action: {
                
            }, label: {
                Image(systemName: "gearshape")
                Text("Settings")
            })
            
            Spacer()
            // Log out
            Button(action: {
                
            }, label: {
                HStack {
                    Image(systemName: "rectangle.portrait.and.arrow.right")
                    Text("Log Out")
                }
                .foregroundColor(.white)
                .padding(10)
                .frame(maxWidth: 150)
            }).background(Color.red)
                .cornerRadius(10)
        }
        .padding(20)
        .font(.title3)
        .foregroundStyle(Color.black).bold()
        .frame(maxWidth: .infinity, alignment: .leading)
    }
}

#Preview {
    SideMenuView()
}
