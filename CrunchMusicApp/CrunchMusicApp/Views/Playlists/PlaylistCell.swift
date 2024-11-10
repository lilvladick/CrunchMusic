import SwiftUI

struct PlaylistCell: View{
    // пока что нет ничего, просто набросок, будет сервер будет и код
    
    var body: some View {
        HStack {
            Image(systemName: "music.note.list")
                .font(.title)
                .padding()
            VStack {
                Text("Playlist name")
                    .font(.headline)
                Text("Playlist description")
            }
            
            Spacer()
            
            Text("Tracks count")
        }
    }
}

#Preview {
    PlaylistCell()
}
