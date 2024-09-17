import SwiftUI

struct SideMenuView: View {
    @Binding var isShowing: Bool
    var body: some View {
        GeometryReader { geometry in
            VStack(alignment: .leading) {
                
                // load track
                Button(action: {
                    
                }, label: {
                    Image(systemName: "square.and.arrow.up")
                    Text("Upload Tracks")
                }).padding(.top, 50)
                
                // your tracks (in view user can pick his tracks only or not)
                Button(action: {
                    
                }, label: {
                    Image(systemName: "music.note.list")
                    Text("Tracks")
                }).padding(.vertical, 15)
                
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
            .foregroundStyle(Color.black).bold()
            .frame(width: geometry.size.width * 0.50,height: geometry.size.height, alignment: .leading)
            .background(Color.white)
            .cornerRadius(20)
            .shadow(radius: 10)
            .offset(x: isShowing ? 0 : -geometry.size.width * 0.50)
            .animation(.easeInOut(duration: 0.3), value: isShowing)
        }
        .edgesIgnoringSafeArea(.all)
    }
}

#Preview {
    @Previewable @State var isShowing: Bool = true
    SideMenuView(isShowing: $isShowing)
}
