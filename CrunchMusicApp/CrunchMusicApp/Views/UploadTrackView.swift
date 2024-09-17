import SwiftUI

struct UploadTrackView: View {
    
    var body: some View {
        NavigationStack {
            VStack {
                
            }
            .toolbar {
                ToolbarItem(placement: .topBarTrailing, content: {
                    Button("Upload") {
                        // upload track button
                    }
                    //.disabled()
                })
            }
            .navigationTitle("Upload Track")
        }
    }
}


#Preview { UploadTrackView() }
